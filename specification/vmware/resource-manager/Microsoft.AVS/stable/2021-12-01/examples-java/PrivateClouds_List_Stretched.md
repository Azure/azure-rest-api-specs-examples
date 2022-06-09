```java
import com.azure.core.util.Context;

/** Samples for PrivateClouds ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/PrivateClouds_List_Stretched.json
     */
    /**
     * Sample code: PrivateClouds_List_Stretched.
     *
     * @param manager Entry point to AvsManager.
     */
    public static void privateCloudsListStretched(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.privateClouds().listByResourceGroup("group1", Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-avs_1.0.0-beta.3/sdk/avs/azure-resourcemanager-avs/README.md) on how to add the SDK to your project and authenticate.
