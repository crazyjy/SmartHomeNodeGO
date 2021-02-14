/**
 * 名称: SmartHomeNode 智能家居控制设备测试
 * 开发板：ESP 8266 MINI D1
 * 时间：2021-02-13
 */
#include <ESP8266WiFi.h>
WiFiClient client;

/***************************** 服务器配置 **********************************/

/* wifi 名称 */
const char *SSID = "OpenWrt";
/* wifi 密码 */
const char *PASSWORD = "12345678";
/* Server 的 IP 地址 */
const char *HOST = "192.168.1.1";
/* Server 的端口号 */
const int PORT = 50001;
/* 节点名称 */
const String NAME = "S1";

/***************************** 控制引脚 *******************************/
/* 控制码 */
String CONTROL = "STATUS/:" + NAME;
const String OPEN = "/:ON";
const String CLOSE = "/:OFF";
/* 接收到的每一位数据 */
char readByte = ' ';
/* 缓存所有接收到的字节 */
String readData = "";
/* 心跳上报标志 */
int heartbeat = 0;
/* LED 灯 */
int LED = 2;

/**
 * 周期上报心跳
 */
void heart() {

    CONTROL = "STATUS/:" + NAME;

    if (digitalRead(LED) == HIGH) {
        CONTROL += CLOSE;
    } else {
        CONTROL += OPEN;
    }

    client.write(&CONTROL[0]);

}

/**
 * TCP 服务端发送过来的数据拼装字符成字符串
 */
void tcpData() {

    while (client.available())
    {

        readByte = client.read();
        readData += readByte;
        delay(5);

    }

}

/**
 * 系统 灯亮一下
 */
void alarmLED() {
    digitalWrite(LED, LOW);
    delay(300);
    digitalWrite(LED, HIGH);
}

/**
 * 开
 */
void switchON() {
    digitalWrite(LED, LOW);
    heart();
}

/**
 * 关
 */
void switchOFF() {
    digitalWrite(LED, HIGH);
    heart();
}

/**
 * 控制开关
 */
void controlSwitch(String constrol) {
    if (constrol == "ON") {
        switchON();
    }
    if (constrol == "OFF") {
        switchOFF();
    }
}

/**
 * 初始化TCP 链接
 */
void linkInit() {

    Serial.begin(115200);
    Serial.print("Connecting to ");
    Serial.println(SSID);
    WiFi.begin(SSID, PASSWORD);

    /* WiFi.status() 返回wifi链接状态  */
    while (WiFi.status() != WL_CONNECTED)
    {
        delay(300);
        Serial.print(".");
    }

    Serial.println("");
    Serial.println("WiFi connected");
    Serial.println("IP address: ");
    Serial.println(WiFi.localIP());

}

/**
 * 设备引脚初始化
 */
void deviceInit() {
    pinMode(LED, OUTPUT);
}

/**
 *  非连接的异常处理
 */
void linkException() {

    while (!client.connected())
    {

        if (!client.connect(HOST, PORT))
        {

            Serial.println("connected excetpicon!");
            alarmLED();

        } else {
            heart();     // 重新连接后登录服务器
        }

        delay(3000);     // 延时3秒再次检测
    }

}

/**
 * 初始化系统
 */
void setup()
{
    /* 初始化 TCP 连接服务 */
    linkInit();
    /* 判断链接状态 */
    linkException();
    /* 设备引脚初始化 */
    deviceInit();
    /* 组合 tcp 发来的数据 */
    tcpData();
    /* 首次登录服务器 */
    heart();
    /* 系统灯点亮一下 */
    alarmLED();
}

/**
 * 系统运行
 */
void loop()
{
    linkException();
    tcpData();

    /* 控制方法 */
    if (readData != "") {
        controlSwitch(readData);
        readData = "";
    }
    /* 心跳间隔上报 */
    if ((++heartbeat) > 300) {
        heartbeat = 0;
        heart();
    }

    delay(100);
}
