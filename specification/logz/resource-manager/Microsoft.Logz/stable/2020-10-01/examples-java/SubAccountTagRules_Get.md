Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-logz_1.0.0-beta.1/sdk/logz/azure-resourcemanager-logz/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for SubAccountTagRules Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/logz/resource-manager/Microsoft.Logz/stable/2020-10-01/examples/SubAccountTagRules_Get.json
     */
    /**
     * Sample code: SubAccountTagRules_Get.
     *
     * @param manager Entry point to LogzManager.
     */
    public static void subAccountTagRulesGet(com.azure.resourcemanager.logz.LogzManager manager) {
        manager
            .subAccountTagRules()
            .getWithResponse("myResourceGroup", "myMonitor", "SubAccount1", "default", Context.NONE);
    }
}
```
