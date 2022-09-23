import com.azure.resourcemanager.security.models.SeverityEnum;
import com.azure.resourcemanager.security.models.SupportedCloudEnum;

/** Samples for CustomAssessmentAutomations Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2021-07-01-preview/examples/CustomAssessmentAutomations/customAssessmentAutomationCreate_example.json
     */
    /**
     * Sample code: Create a Custom Assessment Automation.
     *
     * @param manager Entry point to SecurityManager.
     */
    public static void createACustomAssessmentAutomation(com.azure.resourcemanager.security.SecurityManager manager) {
        manager
            .customAssessmentAutomations()
            .define("MyCustomAssessmentAutomation")
            .withExistingResourceGroup("TestResourceGroup")
            .withCompressedQuery(
                "DQAKAEkAYQBtAF8ARwByAG8AdQBwAA0ACgB8ACAAZQB4AHQAZQBuAGQAIABIAGUAYQBsAHQAaABTAHQAYQB0AHUAcwAgAD0AIABpAGYAZgAoAHQAbwBzAHQAcgBpAG4AZwAoAFIAZQBjAG8AcgBkAC4AVQBzAGUAcgBOAGEAbQBlACkAIABjAG8AbgB0AGEAaQBuAHMAIAAnAHUAcwBlAHIAJwAsACAAJwBVAE4ASABFAEEATABUAEgAWQAnACwAIAAnAEgARQBBAEwAVABIAFkAJwApAA0ACgA=")
            .withSupportedCloud(SupportedCloudEnum.AWS)
            .withSeverity(SeverityEnum.MEDIUM)
            .withDisplayName("Password Policy")
            .withDescription("Data should be encrypted")
            .withRemediationDescription("Encrypt store by...")
            .create();
    }
}
