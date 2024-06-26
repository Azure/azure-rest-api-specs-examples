
import com.azure.resourcemanager.selfhelp.models.CheckNameAvailabilityRequest;

/**
 * Samples for CheckNameAvailability CheckAvailability.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/help/resource-manager/Microsoft.Help/preview/2024-03-01-preview/examples/
     * CheckNameAvailabilityForDiagnosticWhenNameIsNotAvailable.json
     */
    /**
     * Sample code: Example when name is not available for a Diagnostic resource.
     * 
     * @param manager Entry point to SelfHelpManager.
     */
    public static void exampleWhenNameIsNotAvailableForADiagnosticResource(
        com.azure.resourcemanager.selfhelp.SelfHelpManager manager) {
        manager.checkNameAvailabilities().checkAvailabilityWithResponse(
            "subscriptions/0d0fcd2e-c4fd-4349-8497-200edb3923c6",
            new CheckNameAvailabilityRequest().withName("sampleName").withType("Microsoft.Help/diagnostics"),
            com.azure.core.util.Context.NONE);
    }

    /*
     * x-ms-original-file: specification/help/resource-manager/Microsoft.Help/preview/2024-03-01-preview/examples/
     * CheckNameAvailabilityForDiagnosticWhenNameIsAvailable.json
     */
    /**
     * Sample code: Example when name is available for a Diagnostic resource.
     * 
     * @param manager Entry point to SelfHelpManager.
     */
    public static void
        exampleWhenNameIsAvailableForADiagnosticResource(com.azure.resourcemanager.selfhelp.SelfHelpManager manager) {
        manager.checkNameAvailabilities().checkAvailabilityWithResponse(
            "subscriptions/0d0fcd2e-c4fd-4349-8497-200edb3923c6",
            new CheckNameAvailabilityRequest().withName("sampleName").withType("Microsoft.Help/diagnostics"),
            com.azure.core.util.Context.NONE);
    }
}
