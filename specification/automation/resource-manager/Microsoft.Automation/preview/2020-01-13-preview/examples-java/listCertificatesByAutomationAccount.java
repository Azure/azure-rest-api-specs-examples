import com.azure.core.util.Context;

/** Samples for Certificate ListByAutomationAccount. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/listCertificatesByAutomationAccount.json
     */
    /**
     * Sample code: List certificates.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void listCertificates(com.azure.resourcemanager.automation.AutomationManager manager) {
        manager.certificates().listByAutomationAccount("rg", "myAutomationAccount33", Context.NONE);
    }
}
