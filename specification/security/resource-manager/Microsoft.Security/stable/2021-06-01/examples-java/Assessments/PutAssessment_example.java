import com.azure.resourcemanager.security.models.AssessmentStatus;
import com.azure.resourcemanager.security.models.AssessmentStatusCode;
import com.azure.resourcemanager.security.models.AzureResourceDetails;

/** Samples for Assessments CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2021-06-01/examples/Assessments/PutAssessment_example.json
     */
    /**
     * Sample code: Create security recommendation task on a resource.
     *
     * @param manager Entry point to SecurityManager.
     */
    public static void createSecurityRecommendationTaskOnAResource(
        com.azure.resourcemanager.security.SecurityManager manager) {
        manager
            .assessments()
            .define("8bb8be0a-6010-4789-812f-e4d661c4ed0e")
            .withExistingResourceId(
                "subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/myRg/providers/Microsoft.Compute/virtualMachineScaleSets/vmss2")
            .withStatus(new AssessmentStatus().withCode(AssessmentStatusCode.HEALTHY))
            .withResourceDetails(new AzureResourceDetails())
            .create();
    }
}
