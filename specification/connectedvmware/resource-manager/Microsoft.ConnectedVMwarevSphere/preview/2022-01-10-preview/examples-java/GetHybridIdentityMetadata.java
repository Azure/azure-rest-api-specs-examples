import com.azure.core.util.Context;

/** Samples for HybridIdentityMetadata Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/preview/2022-01-10-preview/examples/GetHybridIdentityMetadata.json
     */
    /**
     * Sample code: GetHybridIdentityMetadata.
     *
     * @param manager Entry point to ConnectedVMwareManager.
     */
    public static void getHybridIdentityMetadata(
        com.azure.resourcemanager.connectedvmware.ConnectedVMwareManager manager) {
        manager.hybridIdentityMetadatas().getWithResponse("testrg", "ContosoVm", "default", Context.NONE);
    }
}
