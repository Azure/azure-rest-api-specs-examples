```java
import com.azure.core.util.Context;

/** Samples for SubAccount List. */
public final class Main {
    /*
     * x-ms-original-file: specification/logz/resource-manager/Microsoft.Logz/stable/2020-10-01/examples/SubAccount_ListByResourceGroup.json
     */
    /**
     * Sample code: SubAccount_List.
     *
     * @param manager Entry point to LogzManager.
     */
    public static void subAccountList(com.azure.resourcemanager.logz.LogzManager manager) {
        manager.subAccounts().list("myResourceGroup", "myMonitor", Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-logz_1.0.0-beta.1/sdk/logz/azure-resourcemanager-logz/README.md) on how to add the SDK to your project and authenticate.
