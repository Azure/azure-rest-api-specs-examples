
import com.azure.resourcemanager.workloads.models.SapDatabaseScaleMethod;
import com.azure.resourcemanager.workloads.models.SapDatabaseType;
import com.azure.resourcemanager.workloads.models.SapDeploymentType;
import com.azure.resourcemanager.workloads.models.SapEnvironmentType;
import com.azure.resourcemanager.workloads.models.SapHighAvailabilityType;
import com.azure.resourcemanager.workloads.models.SapProductType;
import com.azure.resourcemanager.workloads.models.SapSizingRecommendationRequest;

/**
 * Samples for ResourceProvider SapSizingRecommendations.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/workloads/resource-manager/Microsoft.Workloads/stable/2023-04-01/examples/sapvirtualinstances/
     * SAPSizingRecommendations_S4HANA_HA_AvZone.json
     */
    /**
     * Sample code: SAPSizingRecommendations_S4HANA_DistributedHA_AvZone.
     * 
     * @param manager Entry point to WorkloadsManager.
     */
    public static void sAPSizingRecommendationsS4HANADistributedHAAvZone(
        com.azure.resourcemanager.workloads.WorkloadsManager manager) {
        manager.resourceProviders().sapSizingRecommendationsWithResponse("centralus",
            new SapSizingRecommendationRequest().withAppLocation("eastus").withEnvironment(SapEnvironmentType.PROD)
                .withSapProduct(SapProductType.S4HANA).withDeploymentType(SapDeploymentType.THREE_TIER).withSaps(75000L)
                .withDbMemory(1024L).withDatabaseType(SapDatabaseType.HANA)
                .withDbScaleMethod(SapDatabaseScaleMethod.SCALE_UP)
                .withHighAvailabilityType(SapHighAvailabilityType.AVAILABILITY_ZONE),
            com.azure.core.util.Context.NONE);
    }
}
