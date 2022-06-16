import com.azure.core.util.Context;
import com.azure.resourcemanager.confidentialledger.models.CheckNameAvailabilityRequest;

/** Samples for ResourceProvider CheckNameAvailability. */
public final class Main {
    /*
     * x-ms-original-file: specification/confidentialledger/resource-manager/Microsoft.ConfidentialLedger/stable/2022-05-13/examples/CheckNameAvailability.json
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
                Context.NONE);
    }
}
