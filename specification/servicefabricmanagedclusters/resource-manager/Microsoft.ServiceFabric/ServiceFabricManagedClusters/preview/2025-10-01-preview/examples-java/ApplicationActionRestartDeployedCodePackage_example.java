
import com.azure.resourcemanager.servicefabricmanagedclusters.models.RestartDeployedCodePackageRequest;

/**
 * Samples for Applications RestartDeployedCodePackage.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/ApplicationActionRestartDeployedCodePackage_example.json
     */
    /**
     * Sample code: Restart deployed code package.
     * 
     * @param manager Entry point to ServiceFabricManagedClustersManager.
     */
    public static void restartDeployedCodePackage(
        com.azure.resourcemanager.servicefabricmanagedclusters.ServiceFabricManagedClustersManager manager) {
        manager.applications().restartDeployedCodePackage("resRg", "myCluster", "myApp",
            new RestartDeployedCodePackageRequest().withNodeName("nt1_0").withServiceManifestName("TestPkg")
                .withCodePackageName("fakeTokenPlaceholder").withCodePackageInstanceId("fakeTokenPlaceholder")
                .withServicePackageActivationId("sharedProcess"),
            com.azure.core.util.Context.NONE);
    }
}
