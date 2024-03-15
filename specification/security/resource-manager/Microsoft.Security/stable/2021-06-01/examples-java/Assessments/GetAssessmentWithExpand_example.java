
import com.azure.resourcemanager.security.models.ExpandEnum;

/**
 * Samples for Assessments Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/stable/2021-06-01/examples/Assessments/
     * GetAssessmentWithExpand_example.json
     */
    /**
     * Sample code: Get security recommendation task from security data location with expand parameter.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void getSecurityRecommendationTaskFromSecurityDataLocationWithExpandParameter(
        com.azure.resourcemanager.security.SecurityManager manager) {
        manager.assessments().getWithResponse(
            "subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/myRg/providers/Microsoft.Compute/virtualMachineScaleSets/vmss2",
            "21300918-b2e3-0346-785f-c77ff57d243b", ExpandEnum.LINKS, com.azure.core.util.Context.NONE);
    }
}
