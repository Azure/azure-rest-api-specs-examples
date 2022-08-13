import com.azure.core.util.Context;
import com.azure.resourcemanager.automation.models.Certificate;

/** Samples for Certificate Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/updateCertificate.json
     */
    /**
     * Sample code: Update a certificate.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void updateACertificate(com.azure.resourcemanager.automation.AutomationManager manager) {
        Certificate resource =
            manager.certificates().getWithResponse("rg", "myAutomationAccount33", "testCert", Context.NONE).getValue();
        resource.update().withName("testCert").withDescription("sample certificate. Description updated").apply();
    }
}
