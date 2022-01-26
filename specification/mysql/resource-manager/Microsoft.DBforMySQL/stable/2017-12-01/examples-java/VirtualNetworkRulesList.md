Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-mysql_1.0.2/sdk/mysql/azure-resourcemanager-mysql/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for VirtualNetworkRules ListByServer. */
public final class Main {
    /*
     * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2017-12-01/examples/VirtualNetworkRulesList.json
     */
    /**
     * Sample code: List virtual network rules.
     *
     * @param manager Entry point to MySqlManager.
     */
    public static void listVirtualNetworkRules(com.azure.resourcemanager.mysql.MySqlManager manager) {
        manager.virtualNetworkRules().listByServer("TestGroup", "vnet-test-svr", Context.NONE);
    }
}
```
