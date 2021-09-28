>此文件为jvm编译器的相关说明
>包括 
>- 命令行工具
>- 搜索class
>- 解析class文件
>- 运行时数据区
>- 指令集和解释器
>- 类与对象

#### 一 命令行工具
##### 1.什么是命令行工具
Java是通过java虚拟机来装载和执行编译文件（class文件）的，java虚拟机通过命令```java option```来启动，```-option```为虚拟机参数，通过这些参数可对虚拟机的运行状态进行调整. 


#### 二 搜索class
##### 1.什么是类路径
Java类路劲告诉java解释器和javac编译器去哪里找到要执行和导入的类。

#### 三 解析class文件
##### 1.理解JAVA Class文件
所有Java文件必须被编译成class文件之后才会被jvm所识别和运用。Java虚拟机中定义的Class文件格式，每一个Class文件都对应唯一一个类或接口的定义信息，但是类或接口不一定都在文件中（**如类和接口可以通过类加载器直接生成**）

每个Class文件都是由8字节为单位的字节流组成，多字节数据项总是按照Big-Endian的顺序进行储存。
Class文件使用类似C语言的伪结构来描述Class文件格式，用于描述类结构的内容定义为项(Item)，表(Table)是由任意数量的可变长度的项组成
```C++
// ClassFile的基本机构
ClassFile { 
    u4 magic; 
    u2 minor_version; 
    u2 major_version; 
    u2 constant_pool_count; 
    cp_info constant_pool[constant_pool_count-1]; 
    u2 access_flags; 
    u2 this_class; 
    u2 super_class; 
    u2 interfaces_count; 
    u2 interfaces[interfaces_count]; 
    u2 fields_count; 
    field_info fields[fields_count]; 
    u2 methods_count; 
    method_info methods[methods_count]; 
    u2 attributes_count; 
    attribute_info attributes[attributes_count]; 
}
```

#### 四 运行时的数据区
这一阶段主要由五个阶段，分别是加载，验证，准备，解析，初始化
##### 1.理解数据区
![avator](https://img-blog.csdn.net/20171014180538873?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvdTAxMTQ2NDUzNg==/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast )
- 方法区
方法区是各个线程共享的内存区域，用于储存被虚拟机加载的类信息，常量，静态常量，即时编译器编译后的代码，运行时常量池。在HotSpot上也被称为"永久代"

##### 2.加载
###### 1.加载过程
- 通过一个类的全限定名来获取定义此类的二进制字节流
- 将这个字节流所代表的静态数据结构转换为方法区的运行时的数据结构(**即三中的classFile类**)
- 在内存中生成一个代表这个类的java.lang.Class对象，作为类的访问入口

#### 五 指令集和解析器

#### 六 类与对象

