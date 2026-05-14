
import com.azure.resourcemanager.fileshares.models.CheckNameAvailabilityRequest;

/**
 * Samples for FileShares CheckNameAvailability.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-01/FileShares_CheckNameAvailability_MaximumSet_Gen.json
     */
    /**
     * Sample code: FileShares_CheckNameAvailability_MaximumSet.
     * 
     * @param manager Entry point to FileSharesManager.
     */
    public static void
        fileSharesCheckNameAvailabilityMaximumSet(com.azure.resourcemanager.fileshares.FileSharesManager manager) {
        manager.fileShares().checkNameAvailabilityWithResponse("westus",
            new CheckNameAvailabilityRequest().withName("fvykqbgmd").withType("Microsoft.FileShares/fileShares"),
            com.azure.core.util.Context.NONE);
    }
}
