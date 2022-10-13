import com.azure.core.util.Context;

/** Samples for HybridIdentityMetadata ListByCluster. */
public final class Main {
    /*
     * x-ms-original-file: specification/hybridaks/resource-manager/Microsoft.HybridContainerService/preview/2022-05-01-preview/examples/HybridIdentityMetadataListByCluster.json
     */
    /**
     * Sample code: HybridIdentityMetadataListByCluster.
     *
     * @param manager Entry point to HybridContainerServiceManager.
     */
    public static void hybridIdentityMetadataListByCluster(
        com.azure.resourcemanager.hybridcontainerservice.HybridContainerServiceManager manager) {
        manager.hybridIdentityMetadatas().listByCluster("testrg", "ContosoTargetCluster", Context.NONE);
    }
}
