
import com.azure.resourcemanager.workloads.models.SapAvailabilityZoneDetailsRequest;
import com.azure.resourcemanager.workloads.models.SapDatabaseType;
import com.azure.resourcemanager.workloads.models.SapProductType;

/**
 * Samples for ResourceProvider SapAvailabilityZoneDetails.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/workloads/resource-manager/Microsoft.Workloads/stable/2023-04-01/examples/sapvirtualinstances/
     * SAPAvailabilityZoneDetails_northeurope.json
     */
    /**
     * Sample code: SAPAvailabilityZoneDetails_northeurope.
     * 
     * @param manager Entry point to WorkloadsManager.
     */
    public static void
        sAPAvailabilityZoneDetailsNortheurope(com.azure.resourcemanager.workloads.WorkloadsManager manager) {
        manager.resourceProviders().sapAvailabilityZoneDetailsWithResponse(
            "centralus", new SapAvailabilityZoneDetailsRequest().withAppLocation("northeurope")
                .withSapProduct(SapProductType.S4HANA).withDatabaseType(SapDatabaseType.HANA),
            com.azure.core.util.Context.NONE);
    }
}
