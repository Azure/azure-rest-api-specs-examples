
import com.azure.resourcemanager.workloadssapvirtualinstance.models.SapDatabaseScaleMethod;
import com.azure.resourcemanager.workloadssapvirtualinstance.models.SapDatabaseType;
import com.azure.resourcemanager.workloadssapvirtualinstance.models.SapDeploymentType;
import com.azure.resourcemanager.workloadssapvirtualinstance.models.SapEnvironmentType;
import com.azure.resourcemanager.workloadssapvirtualinstance.models.SapProductType;
import com.azure.resourcemanager.workloadssapvirtualinstance.models.SapSizingRecommendationRequest;

/**
 * Samples for SapVirtualInstances GetSizingRecommendations.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01/SapVirtualInstances_InvokeSizingRecommendations_S4HANA_Distributed.json
     */
    /**
     * Sample code: SAP sizing recommendations for non HA distributed system.
     * 
     * @param manager Entry point to WorkloadsSapVirtualInstanceManager.
     */
    public static void sAPSizingRecommendationsForNonHADistributedSystem(
        com.azure.resourcemanager.workloadssapvirtualinstance.WorkloadsSapVirtualInstanceManager manager) {
        manager.sapVirtualInstances().getSizingRecommendationsWithResponse("centralus",
            new SapSizingRecommendationRequest().withAppLocation("eastus").withEnvironment(SapEnvironmentType.PROD)
                .withSapProduct(SapProductType.S4HANA).withDeploymentType(SapDeploymentType.THREE_TIER).withSaps(20000L)
                .withDbMemory(1024L).withDatabaseType(SapDatabaseType.HANA)
                .withDbScaleMethod(SapDatabaseScaleMethod.SCALE_UP),
            com.azure.core.util.Context.NONE);
    }
}
