/** Samples for HybridIdentityMetadata Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/hybridaks/resource-manager/Microsoft.HybridContainerService/preview/2022-09-01-preview/examples/GetHybridIdentityMetadata.json
     */
    /**
     * Sample code: GetHybridIdentityMetadata.
     *
     * @param manager Entry point to HybridContainerServiceManager.
     */
    public static void getHybridIdentityMetadata(
        com.azure.resourcemanager.hybridcontainerservice.HybridContainerServiceManager manager) {
        manager
            .hybridIdentityMetadatas()
            .getWithResponse("testrg", "ContosoTargetCluster", "default", com.azure.core.util.Context.NONE);
    }
}
