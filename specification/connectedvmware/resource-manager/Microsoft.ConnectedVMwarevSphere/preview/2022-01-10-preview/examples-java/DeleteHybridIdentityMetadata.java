import com.azure.core.util.Context;

/** Samples for HybridIdentityMetadata Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/preview/2022-01-10-preview/examples/DeleteHybridIdentityMetadata.json
     */
    /**
     * Sample code: DeleteHybridIdentityMetadata.
     *
     * @param manager Entry point to ConnectedVMwareManager.
     */
    public static void deleteHybridIdentityMetadata(
        com.azure.resourcemanager.connectedvmware.ConnectedVMwareManager manager) {
        manager.hybridIdentityMetadatas().deleteWithResponse("testrg", "ContosoVm", "default", Context.NONE);
    }
}
