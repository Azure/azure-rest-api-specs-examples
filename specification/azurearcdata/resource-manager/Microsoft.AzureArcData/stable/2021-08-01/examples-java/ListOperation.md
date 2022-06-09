```java
import com.azure.core.util.Context;

/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/azurearcdata/resource-manager/Microsoft.AzureArcData/stable/2021-08-01/examples/ListOperation.json
     */
    /**
     * Sample code: Lists all of the available Azure Data Services on Azure Arc API operations.
     *
     * @param manager Entry point to AzureArcDataManager.
     */
    public static void listsAllOfTheAvailableAzureDataServicesOnAzureArcAPIOperations(
        com.azure.resourcemanager.azurearcdata.AzureArcDataManager manager) {
        manager.operations().list(Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-azurearcdata_1.0.0-beta.2/sdk/azurearcdata/azure-resourcemanager-azurearcdata/README.md) on how to add the SDK to your project and authenticate.
