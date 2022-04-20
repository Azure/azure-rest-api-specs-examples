Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-redisenterprise_1.1.0-beta.1/sdk/redisenterprise/azure-resourcemanager-redisenterprise/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/redisenterprise/resource-manager/Microsoft.Cache/stable/2022-01-01/examples/OperationsList.json
     */
    /**
     * Sample code: OperationsList.
     *
     * @param manager Entry point to RedisEnterpriseManager.
     */
    public static void operationsList(com.azure.resourcemanager.redisenterprise.RedisEnterpriseManager manager) {
        manager.operations().list(Context.NONE);
    }
}
```
