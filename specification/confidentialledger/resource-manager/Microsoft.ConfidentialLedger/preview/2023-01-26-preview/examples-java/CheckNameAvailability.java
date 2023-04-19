import com.azure.resourcemanager.confidentialledger.models.CheckNameAvailabilityRequest;

/** Samples for ResourceProvider CheckNameAvailability. */
public final class Main {
    /*
     * x-ms-original-file: specification/confidentialledger/resource-manager/Microsoft.ConfidentialLedger/preview/2023-01-26-preview/examples/CheckNameAvailability.json
     */
    /**
     * Sample code: CheckNameAvailability.
     *
     * @param manager Entry point to ConfidentialLedgerManager.
     */
    public static void checkNameAvailability(
        com.azure.resourcemanager.confidentialledger.ConfidentialLedgerManager manager) {
        manager
            .resourceProviders()
            .checkNameAvailabilityWithResponse(
                new CheckNameAvailabilityRequest()
                    .withName("sample-name")
                    .withType("Microsoft.ConfidentialLedger/ledgers"),
                com.azure.core.util.Context.NONE);
    }
}
