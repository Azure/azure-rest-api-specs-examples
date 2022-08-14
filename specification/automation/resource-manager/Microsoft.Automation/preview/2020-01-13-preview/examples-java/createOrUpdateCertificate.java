/** Samples for Certificate CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/createOrUpdateCertificate.json
     */
    /**
     * Sample code: Create or update a certificate.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void createOrUpdateACertificate(com.azure.resourcemanager.automation.AutomationManager manager) {
        manager
            .certificates()
            .define("testCert")
            .withExistingAutomationAccount("rg", "myAutomationAccount18")
            .withName("testCert")
            .withBase64Value("base 64 value of cert")
            .withDescription("Sample Cert")
            .withThumbprint("thumbprint of cert")
            .withIsExportable(false)
            .create();
    }
}
