import com.azure.core.util.Context;

/** Samples for Certificate Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/deleteCertificate.json
     */
    /**
     * Sample code: Delete a certificate.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void deleteACertificate(com.azure.resourcemanager.automation.AutomationManager manager) {
        manager.certificates().deleteWithResponse("rg", "myAutomationAccount33", "testCert", Context.NONE);
    }
}
