```java
import com.azure.core.util.Context;

/** Samples for SqlManagedInstances ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/azurearcdata/resource-manager/Microsoft.AzureArcData/stable/2021-08-01/examples/ListByResourceGroupSqlManagedInstance.json
     */
    /**
     * Sample code: Gets all SQL Instance in a resource group.
     *
     * @param manager Entry point to AzureArcDataManager.
     */
    public static void getsAllSQLInstanceInAResourceGroup(
        com.azure.resourcemanager.azurearcdata.AzureArcDataManager manager) {
        manager.sqlManagedInstances().listByResourceGroup("testrg", Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-azurearcdata_1.0.0-beta.2/sdk/azurearcdata/azure-resourcemanager-azurearcdata/README.md) on how to add the SDK to your project and authenticate.
