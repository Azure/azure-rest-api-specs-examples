
import com.azure.resourcemanager.servicefabric.models.UpgradableVersionsDescription;

/**
 * Samples for Clusters ListUpgradableVersions.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/ListUpgradableVersionsPath_example.json
     */
    /**
     * Sample code: Get upgrade path.
     * 
     * @param manager Entry point to ServiceFabricManager.
     */
    public static void getUpgradePath(com.azure.resourcemanager.servicefabric.ServiceFabricManager manager) {
        manager.clusters().listUpgradableVersionsWithResponse("resRg", "myCluster",
            new UpgradableVersionsDescription().withTargetVersion("7.2.432.9590"), com.azure.core.util.Context.NONE);
    }
}
