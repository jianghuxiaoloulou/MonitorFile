# 进程监控系统

# 进程监控系统，定时判断进程是否存在，不存在，启动进程
# 通过读取进程配置文件获取需要监控的进程
# 监控进程配置文件格式：进程名+ " " + 进程路径

github:


Cron表达式范例：
    每隔5秒执行一次：*/5 * * * * ?

    每隔1分钟执行一次：0 */1 * * * ?

    每天23点执行一次：0 0 23 * * ?

    每天凌晨1点执行一次：0 0 1 * * ?

    每月1号凌晨1点执行一次：0 0 1 1 * ?

    每月最后一天23点执行一次：0 0 23 L * ?

    每周星期天凌晨1点实行一次：0 0 1 ? * L

    在26分、29分、33分执行一次：0 26,29,33 * * * ?

    每天的0点、13点、18点、21点都执行一次：0 0 0,13,18,21 * * ?