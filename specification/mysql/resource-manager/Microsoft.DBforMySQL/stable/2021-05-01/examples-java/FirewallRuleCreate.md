Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-mysqlflexibleserver_1.0.0-beta.2/sdk/mysqlflexibleserver/azure-resourcemanager-mysqlflexibleserver/README.md) on how to add the SDK to your project and authenticate.

```java
/** Samples for FirewallRules CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2021-05-01/examples/FirewallRuleCreate.json
     */
    /**
     * Sample code: Create a firewall rule.
     *
     * @param manager Entry point to MySqlManager.
     */
    public static void createAFirewallRule(com.azure.resourcemanager.mysqlflexibleserver.MySqlManager manager) {
        manager
            .firewallRules()
            .define("rule1")
            .withExistingFlexibleServer("TestGroup", "testserver")
            .withStartIpAddress("0.0.0.0")
            .withEndIpAddress("255.255.255.255")
            .create();
    }
}
```
