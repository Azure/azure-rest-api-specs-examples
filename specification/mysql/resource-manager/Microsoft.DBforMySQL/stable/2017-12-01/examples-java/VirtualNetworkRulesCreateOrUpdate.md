Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-mysql_1.0.2/sdk/mysql/azure-resourcemanager-mysql/README.md) on how to add the SDK to your project and authenticate.

```java
/** Samples for VirtualNetworkRules CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2017-12-01/examples/VirtualNetworkRulesCreateOrUpdate.json
     */
    /**
     * Sample code: Create or update a virtual network rule.
     *
     * @param manager Entry point to MySqlManager.
     */
    public static void createOrUpdateAVirtualNetworkRule(com.azure.resourcemanager.mysql.MySqlManager manager) {
        manager
            .virtualNetworkRules()
            .define("vnet-firewall-rule")
            .withExistingServer("TestGroup", "vnet-test-svr")
            .withVirtualNetworkSubnetId(
                "/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestGroup/providers/Microsoft.Network/virtualNetworks/testvnet/subnets/testsubnet")
            .withIgnoreMissingVnetServiceEndpoint(false)
            .create();
    }
}
```
