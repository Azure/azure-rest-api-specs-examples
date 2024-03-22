
import com.azure.resourcemanager.workloadssapvirtualinstance.models.SapDatabaseScaleMethod;
import com.azure.resourcemanager.workloadssapvirtualinstance.models.SapDatabaseType;
import com.azure.resourcemanager.workloadssapvirtualinstance.models.SapDeploymentType;
import com.azure.resourcemanager.workloadssapvirtualinstance.models.SapEnvironmentType;
import com.azure.resourcemanager.workloadssapvirtualinstance.models.SapProductType;
import com.azure.resourcemanager.workloadssapvirtualinstance.models.SapSizingRecommendationRequest;

/**
 * Samples for ResourceProvider SapSizingRecommendations.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/workloads/resource-manager/Microsoft.Workloads/SAPVirtualInstance/preview/2023-10-01-preview/
     * examples/sapvirtualinstances/SAPSizingRecommendations_S4HANA_SingleServer.json
     */
    /**
     * Sample code: SAPSizingRecommendations_S4HANA_SingleServer.
     * 
     * @param manager Entry point to WorkloadsSapVirtualInstanceManager.
     */
    public static void sAPSizingRecommendationsS4HANASingleServer(
        com.azure.resourcemanager.workloadssapvirtualinstance.WorkloadsSapVirtualInstanceManager manager) {
        manager.resourceProviders().sapSizingRecommendationsWithResponse("centralus",
            new SapSizingRecommendationRequest().withAppLocation("eastus").withEnvironment(SapEnvironmentType.NON_PROD)
                .withSapProduct(SapProductType.S4HANA).withDeploymentType(SapDeploymentType.SINGLE_SERVER)
                .withSaps(60000L).withDbMemory(2000L).withDatabaseType(SapDatabaseType.HANA)
                .withDbScaleMethod(SapDatabaseScaleMethod.SCALE_UP),
            com.azure.core.util.Context.NONE);
    }
}
