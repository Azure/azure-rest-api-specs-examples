```java
import com.azure.core.util.Context;

/** Samples for Labs ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/labservices/resource-manager/Microsoft.LabServices/preview/2021-11-15-preview/examples/Labs/listResourceGroupLabs.json
     */
    /**
     * Sample code: listResourceGroupLabs.
     *
     * @param manager Entry point to LabServicesManager.
     */
    public static void listResourceGroupLabs(com.azure.resourcemanager.labservices.LabServicesManager manager) {
        manager.labs().listByResourceGroup("testrg123", Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-labservices_1.0.0-beta.2/sdk/labservices/azure-resourcemanager-labservices/README.md) on how to add the SDK to your project and authenticate.
