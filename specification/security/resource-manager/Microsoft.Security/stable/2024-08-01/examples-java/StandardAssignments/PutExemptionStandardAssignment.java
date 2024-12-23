
import com.azure.resourcemanager.security.models.AssignedAssessmentItem;
import com.azure.resourcemanager.security.models.AssignedStandardItem;
import com.azure.resourcemanager.security.models.Effect;
import com.azure.resourcemanager.security.models.ExemptionCategory;
import com.azure.resourcemanager.security.models.StandardAssignmentPropertiesExemptionData;
import java.time.OffsetDateTime;

/**
 * Samples for StandardAssignments Create.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/stable/2024-08-01/examples/StandardAssignments/
     * PutExemptionStandardAssignment.json
     */
    /**
     * Sample code: Put exemption standard assignment.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void putExemptionStandardAssignment(com.azure.resourcemanager.security.SecurityManager manager) {
        manager.standardAssignments().define("1f3afdf9-d0c9-4c3d-847f-89da613e70a8").withExistingResourceId(
            "subscriptions/212f9889-769e-45ae-ab43-6da33674bd26/resourceGroups/ANAT_TEST_RG/providers/Microsoft.Compute/virtualMachines/anatTestE2LA")
            .withDisplayName("Test exemption").withDescription("Exemption description")
            .withAssignedStandard(new AssignedStandardItem()
                .withId("/providers/Microsoft.Security/securityStandards/1f3afdf9-d0c9-4c3d-847f-89da613e70a8"))
            .withEffect(Effect.EXEMPT).withExpiresOn(OffsetDateTime.parse("2022-05-01T19:50:47.083633Z"))
            .withExemptionData(
                new StandardAssignmentPropertiesExemptionData().withExemptionCategory(ExemptionCategory.WAIVER)
                    .withAssignedAssessment(new AssignedAssessmentItem().withAssessmentKey("fakeTokenPlaceholder")))
            .create();
    }
}
