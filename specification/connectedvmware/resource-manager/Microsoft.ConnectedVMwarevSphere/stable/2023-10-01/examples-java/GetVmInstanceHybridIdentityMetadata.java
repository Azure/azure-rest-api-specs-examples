/** Samples for VmInstanceHybridIdentityMetadata Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/stable/2023-10-01/examples/GetVmInstanceHybridIdentityMetadata.json
     */
    /**
     * Sample code: GetHybridIdentityMetadata.
     *
     * @param manager Entry point to ConnectedVMwareManager.
     */
    public static void getHybridIdentityMetadata(
        com.azure.resourcemanager.connectedvmware.ConnectedVMwareManager manager) {
        manager
            .vmInstanceHybridIdentityMetadatas()
            .getWithResponse(
                "subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.HybridCompute/machines/DemoVM",
                com.azure.core.util.Context.NONE);
    }
}
