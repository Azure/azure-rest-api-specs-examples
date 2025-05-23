
import com.azure.resourcemanager.security.models.AssessmentType;
import com.azure.resourcemanager.security.models.Categories;
import com.azure.resourcemanager.security.models.ImplementationEffort;
import com.azure.resourcemanager.security.models.Severity;
import com.azure.resourcemanager.security.models.Threats;
import com.azure.resourcemanager.security.models.UserImpact;
import java.util.Arrays;

/**
 * Samples for AssessmentsMetadata CreateInSubscription.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/stable/2021-06-01/examples/AssessmentsMetadata/
     * CreateAssessmentsMetadata_subscription_example.json
     */
    /**
     * Sample code: Create security assessment metadata for subscription.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void
        createSecurityAssessmentMetadataForSubscription(com.azure.resourcemanager.security.SecurityManager manager) {
        manager.assessmentsMetadatas().define("ca039e75-a276-4175-aebc-bcd41e4b14b7")
            .withDisplayName("Install endpoint protection solution on virtual machine scale sets")
            .withDescription(
                "Install an endpoint protection solution on your virtual machines scale sets, to protect them from threats and vulnerabilities.")
            .withRemediationDescription(
                "To install an endpoint protection solution: 1.  <a href=\"https://docs.microsoft.com/azure/virtual-machine-scale-sets/virtual-machine-scale-sets-faq#how-do-i-turn-on-antimalware-in-my-virtual-machine-scale-set\">Follow the instructions in How do I turn on antimalware in my virtual machine scale set</a>")
            .withCategories(Arrays.asList(Categories.COMPUTE)).withSeverity(Severity.MEDIUM)
            .withUserImpact(UserImpact.LOW).withImplementationEffort(ImplementationEffort.LOW)
            .withThreats(Arrays.asList(Threats.DATA_EXFILTRATION, Threats.DATA_SPILLAGE, Threats.MALICIOUS_INSIDER))
            .withAssessmentType(AssessmentType.CUSTOMER_MANAGED).create();
    }
}
