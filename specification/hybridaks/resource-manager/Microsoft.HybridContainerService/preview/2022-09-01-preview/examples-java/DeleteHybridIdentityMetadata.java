/** Samples for HybridIdentityMetadata Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/hybridaks/resource-manager/Microsoft.HybridContainerService/preview/2022-09-01-preview/examples/DeleteHybridIdentityMetadata.json
     */
    /**
     * Sample code: DeleteHybridIdentityMetadata.
     *
     * @param manager Entry point to HybridContainerServiceManager.
     */
    public static void deleteHybridIdentityMetadata(
        com.azure.resourcemanager.hybridcontainerservice.HybridContainerServiceManager manager) {
        manager
            .hybridIdentityMetadatas()
            .deleteWithResponse("testrg", "ContosoTargetCluster", "default", com.azure.core.util.Context.NONE);
    }
}
