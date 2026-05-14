
import com.azure.resourcemanager.fileshares.models.CheckNameAvailabilityRequest;

/**
 * Samples for FileShares CheckNameAvailability.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-01/FileShares_CheckNameAvailability_MinimumSet_Gen.json
     */
    /**
     * Sample code: FileShares_CheckNameAvailability_MinimumSet.
     * 
     * @param manager Entry point to FileSharesManager.
     */
    public static void
        fileSharesCheckNameAvailabilityMinimumSet(com.azure.resourcemanager.fileshares.FileSharesManager manager) {
        manager.fileShares().checkNameAvailabilityWithResponse("westus", new CheckNameAvailabilityRequest(),
            com.azure.core.util.Context.NONE);
    }
}
