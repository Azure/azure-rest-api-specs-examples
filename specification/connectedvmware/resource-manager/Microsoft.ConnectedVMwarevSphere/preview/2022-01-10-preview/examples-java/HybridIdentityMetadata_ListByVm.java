import com.azure.core.util.Context;

/** Samples for HybridIdentityMetadata ListByVm. */
public final class Main {
    /*
     * x-ms-original-file: specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/preview/2022-01-10-preview/examples/HybridIdentityMetadata_ListByVm.json
     */
    /**
     * Sample code: HybridIdentityMetadataListByVm.
     *
     * @param manager Entry point to ConnectedVMwareManager.
     */
    public static void hybridIdentityMetadataListByVm(
        com.azure.resourcemanager.connectedvmware.ConnectedVMwareManager manager) {
        manager.hybridIdentityMetadatas().listByVm("testrg", "ContosoVm", Context.NONE);
    }
}
