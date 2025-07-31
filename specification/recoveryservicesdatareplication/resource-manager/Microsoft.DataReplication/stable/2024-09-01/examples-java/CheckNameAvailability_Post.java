
import com.azure.resourcemanager.recoveryservicesdatareplication.models.CheckNameAvailabilityModel;

/**
 * Samples for CheckNameAvailability Post.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01/CheckNameAvailability_Post.json
     */
    /**
     * Sample code: Performs the resource name availability check.
     * 
     * @param manager Entry point to RecoveryServicesDataReplicationManager.
     */
    public static void performsTheResourceNameAvailabilityCheck(
        com.azure.resourcemanager.recoveryservicesdatareplication.RecoveryServicesDataReplicationManager manager) {
        manager.checkNameAvailabilities().postWithResponse("trfqtbtmusswpibw",
            new CheckNameAvailabilityModel().withName("updkdcixs").withType("gngmcancdauwhdixjjvqnfkvqc"),
            com.azure.core.util.Context.NONE);
    }
}
