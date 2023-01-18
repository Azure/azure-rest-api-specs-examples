import com.azure.resourcemanager.purview.models.CheckNameAvailabilityRequest;

/** Samples for Accounts CheckNameAvailability. */
public final class Main {
    /*
     * x-ms-original-file: specification/purview/resource-manager/Microsoft.Purview/stable/2021-07-01/examples/Accounts_CheckNameAvailability.json
     */
    /**
     * Sample code: Accounts_CheckNameAvailability.
     *
     * @param manager Entry point to PurviewManager.
     */
    public static void accountsCheckNameAvailability(com.azure.resourcemanager.purview.PurviewManager manager) {
        manager
            .accounts()
            .checkNameAvailabilityWithResponse(
                new CheckNameAvailabilityRequest().withName("account1").withType("Microsoft.Purview/accounts"),
                com.azure.core.util.Context.NONE);
    }
}
