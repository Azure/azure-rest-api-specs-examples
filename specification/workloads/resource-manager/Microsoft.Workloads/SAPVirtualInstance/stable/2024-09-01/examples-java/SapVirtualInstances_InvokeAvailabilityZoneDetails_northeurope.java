
import com.azure.resourcemanager.workloadssapvirtualinstance.models.SapAvailabilityZoneDetailsRequest;
import com.azure.resourcemanager.workloadssapvirtualinstance.models.SapDatabaseType;
import com.azure.resourcemanager.workloadssapvirtualinstance.models.SapProductType;

/**
 * Samples for SapVirtualInstances GetAvailabilityZoneDetails.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01/SapVirtualInstances_InvokeAvailabilityZoneDetails_northeurope.json
     */
    /**
     * Sample code: SAP Availability zone details in north europe.
     * 
     * @param manager Entry point to WorkloadsSapVirtualInstanceManager.
     */
    public static void sAPAvailabilityZoneDetailsInNorthEurope(
        com.azure.resourcemanager.workloadssapvirtualinstance.WorkloadsSapVirtualInstanceManager manager) {
        manager.sapVirtualInstances().getAvailabilityZoneDetailsWithResponse(
            "northeurope", new SapAvailabilityZoneDetailsRequest().withAppLocation("northeurope")
                .withSapProduct(SapProductType.S4HANA).withDatabaseType(SapDatabaseType.HANA),
            com.azure.core.util.Context.NONE);
    }
}
