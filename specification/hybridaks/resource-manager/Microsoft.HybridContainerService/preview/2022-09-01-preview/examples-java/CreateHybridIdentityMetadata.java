/** Samples for HybridIdentityMetadata Put. */
public final class Main {
    /*
     * x-ms-original-file: specification/hybridaks/resource-manager/Microsoft.HybridContainerService/preview/2022-09-01-preview/examples/CreateHybridIdentityMetadata.json
     */
    /**
     * Sample code: CreateHybridIdentityMetadata.
     *
     * @param manager Entry point to HybridContainerServiceManager.
     */
    public static void createHybridIdentityMetadata(
        com.azure.resourcemanager.hybridcontainerservice.HybridContainerServiceManager manager) {
        manager
            .hybridIdentityMetadatas()
            .define("default")
            .withExistingProvisionedCluster("testrg", "ContosoTargetCluster")
            .withResourceUid("f8b82dff-38ef-4220-99ef-d3a3f86ddc6c")
            .withPublicKey("8ec7d60c-9700-40b1-8e6e-e5b2f6f477f2")
            .create();
    }
}
