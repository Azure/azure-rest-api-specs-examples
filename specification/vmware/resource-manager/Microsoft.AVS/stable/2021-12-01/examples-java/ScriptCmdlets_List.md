```java
import com.azure.core.util.Context;

/** Samples for ScriptCmdlets List. */
public final class Main {
    /*
     * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/ScriptCmdlets_List.json
     */
    /**
     * Sample code: ScriptCmdlets_List.
     *
     * @param manager Entry point to AvsManager.
     */
    public static void scriptCmdletsList(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.scriptCmdlets().list("group1", "{privateCloudName}", "{scriptPackageName}", Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-avs_1.0.0-beta.3/sdk/avs/azure-resourcemanager-avs/README.md) on how to add the SDK to your project and authenticate.
