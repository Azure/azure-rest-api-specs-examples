```java
import com.azure.core.util.Context;

/** Samples for VmmServers ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/scvmm/resource-manager/Microsoft.ScVmm/preview/2020-06-05-preview/examples/ListVMMServersByResourceGroup.json
     */
    /**
     * Sample code: ListVmmServersByResourceGroup.
     *
     * @param manager Entry point to ScvmmManager.
     */
    public static void listVmmServersByResourceGroup(com.azure.resourcemanager.scvmm.ScvmmManager manager) {
        manager.vmmServers().listByResourceGroup("testrg", Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-scvmm_1.0.0-beta.1/sdk/scvmm/azure-resourcemanager-scvmm/README.md) on how to add the SDK to your project and authenticate.
