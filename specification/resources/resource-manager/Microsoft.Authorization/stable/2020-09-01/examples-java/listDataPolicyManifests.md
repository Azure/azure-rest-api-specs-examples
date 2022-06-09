```java
import com.azure.core.util.Context;

/** Samples for DataPolicyManifests List. */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2020-09-01/examples/listDataPolicyManifests.json
     */
    /**
     * Sample code: List data policy manifests.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listDataPolicyManifests(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.genericResources().manager().policyClient().getDataPolicyManifests().list(null, Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.15.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.
