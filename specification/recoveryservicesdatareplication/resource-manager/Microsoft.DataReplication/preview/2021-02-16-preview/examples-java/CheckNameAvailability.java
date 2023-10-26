import com.azure.resourcemanager.recoveryservicesdatareplication.models.CheckNameAvailabilityModel;

/** Samples for ResourceProvider CheckNameAvailability. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesdatareplication/resource-manager/Microsoft.DataReplication/preview/2021-02-16-preview/examples/CheckNameAvailability.json
     */
    /**
     * Sample code: CheckNameAvailability.
     *
     * @param manager Entry point to RecoveryServicesDataReplicationManager.
     */
    public static void checkNameAvailability(
        com.azure.resourcemanager.recoveryservicesdatareplication.RecoveryServicesDataReplicationManager manager) {
        manager
            .resourceProviders()
            .checkNameAvailabilityWithResponse(
                "trfqtbtmusswpibw",
                new CheckNameAvailabilityModel().withName("updkdcixs").withType("gngmcancdauwhdixjjvqnfkvqc"),
                com.azure.core.util.Context.NONE);
    }
}
