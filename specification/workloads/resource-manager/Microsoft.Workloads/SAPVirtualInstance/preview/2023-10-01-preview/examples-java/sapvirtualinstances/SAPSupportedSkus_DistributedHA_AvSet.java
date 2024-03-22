
import com.azure.resourcemanager.workloadssapvirtualinstance.models.SapDatabaseType;
import com.azure.resourcemanager.workloadssapvirtualinstance.models.SapDeploymentType;
import com.azure.resourcemanager.workloadssapvirtualinstance.models.SapEnvironmentType;
import com.azure.resourcemanager.workloadssapvirtualinstance.models.SapHighAvailabilityType;
import com.azure.resourcemanager.workloadssapvirtualinstance.models.SapProductType;
import com.azure.resourcemanager.workloadssapvirtualinstance.models.SapSupportedSkusRequest;

/**
 * Samples for ResourceProvider SapSupportedSku.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/workloads/resource-manager/Microsoft.Workloads/SAPVirtualInstance/preview/2023-10-01-preview/
     * examples/sapvirtualinstances/SAPSupportedSkus_DistributedHA_AvSet.json
     */
    /**
     * Sample code: SAPSupportedSkus_DistributedHA_AvSet.
     * 
     * @param manager Entry point to WorkloadsSapVirtualInstanceManager.
     */
    public static void sAPSupportedSkusDistributedHAAvSet(
        com.azure.resourcemanager.workloadssapvirtualinstance.WorkloadsSapVirtualInstanceManager manager) {
        manager.resourceProviders().sapSupportedSkuWithResponse("centralus",
            new SapSupportedSkusRequest().withAppLocation("eastus").withEnvironment(SapEnvironmentType.PROD)
                .withSapProduct(SapProductType.S4HANA).withDeploymentType(SapDeploymentType.THREE_TIER)
                .withDatabaseType(SapDatabaseType.HANA)
                .withHighAvailabilityType(SapHighAvailabilityType.AVAILABILITY_SET),
            com.azure.core.util.Context.NONE);
    }
}
