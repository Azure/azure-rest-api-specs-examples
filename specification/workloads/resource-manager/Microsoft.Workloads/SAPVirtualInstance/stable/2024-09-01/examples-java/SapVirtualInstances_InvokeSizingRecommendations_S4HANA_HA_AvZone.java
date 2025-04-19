
import com.azure.resourcemanager.workloadssapvirtualinstance.models.SapDatabaseScaleMethod;
import com.azure.resourcemanager.workloadssapvirtualinstance.models.SapDatabaseType;
import com.azure.resourcemanager.workloadssapvirtualinstance.models.SapDeploymentType;
import com.azure.resourcemanager.workloadssapvirtualinstance.models.SapEnvironmentType;
import com.azure.resourcemanager.workloadssapvirtualinstance.models.SapHighAvailabilityType;
import com.azure.resourcemanager.workloadssapvirtualinstance.models.SapProductType;
import com.azure.resourcemanager.workloadssapvirtualinstance.models.SapSizingRecommendationRequest;

/**
 * Samples for SapVirtualInstances GetSizingRecommendations.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01/SapVirtualInstances_InvokeSizingRecommendations_S4HANA_HA_AvZone.json
     */
    /**
     * Sample code: SAP sizing recommendations for HA with availability zone.
     * 
     * @param manager Entry point to WorkloadsSapVirtualInstanceManager.
     */
    public static void sAPSizingRecommendationsForHAWithAvailabilityZone(
        com.azure.resourcemanager.workloadssapvirtualinstance.WorkloadsSapVirtualInstanceManager manager) {
        manager.sapVirtualInstances().getSizingRecommendationsWithResponse("centralus",
            new SapSizingRecommendationRequest().withAppLocation("eastus").withEnvironment(SapEnvironmentType.PROD)
                .withSapProduct(SapProductType.S4HANA).withDeploymentType(SapDeploymentType.THREE_TIER).withSaps(75000L)
                .withDbMemory(1024L).withDatabaseType(SapDatabaseType.HANA)
                .withDbScaleMethod(SapDatabaseScaleMethod.SCALE_UP)
                .withHighAvailabilityType(SapHighAvailabilityType.AVAILABILITY_ZONE),
            com.azure.core.util.Context.NONE);
    }
}
