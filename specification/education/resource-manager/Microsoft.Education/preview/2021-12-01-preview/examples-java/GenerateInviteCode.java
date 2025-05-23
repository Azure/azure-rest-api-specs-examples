
import com.azure.resourcemanager.education.models.InviteCodeGenerateRequest;

/**
 * Samples for Labs GenerateInviteCode.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/education/resource-manager/Microsoft.Education/preview/2021-12-01-preview/examples/
     * GenerateInviteCode.json
     */
    /**
     * Sample code: CreateLab.
     * 
     * @param manager Entry point to EducationManager.
     */
    public static void createLab(com.azure.resourcemanager.education.EducationManager manager) {
        manager.labs().generateInviteCodeWithResponse("{billingAccountName}", "{billingProfileName}",
            "{invoiceSectionName}", new InviteCodeGenerateRequest().withMaxStudentCount(10.0F), null,
            com.azure.core.util.Context.NONE);
    }
}
