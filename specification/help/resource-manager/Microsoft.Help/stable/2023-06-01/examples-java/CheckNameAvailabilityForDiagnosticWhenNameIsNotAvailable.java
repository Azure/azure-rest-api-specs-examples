import com.azure.resourcemanager.selfhelp.models.CheckNameAvailabilityRequest;

/** Samples for Diagnostics CheckNameAvailability. */
public final class Main {
    /*
     * x-ms-original-file: specification/help/resource-manager/Microsoft.Help/stable/2023-06-01/examples/CheckNameAvailabilityForDiagnosticWhenNameIsNotAvailable.json
     */
    /**
     * Sample code: Example when name is not available for a Diagnostic resource.
     *
     * @param manager Entry point to SelfHelpManager.
     */
    public static void exampleWhenNameIsNotAvailableForADiagnosticResource(
        com.azure.resourcemanager.selfhelp.SelfHelpManager manager) {
        manager
            .diagnostics()
            .checkNameAvailabilityWithResponse(
                "subscriptions/0d0fcd2e-c4fd-4349-8497-200edb3923c6",
                new CheckNameAvailabilityRequest().withName("sampleName").withType("Microsoft.Help/diagnostics"),
                com.azure.core.util.Context.NONE);
    }
}
