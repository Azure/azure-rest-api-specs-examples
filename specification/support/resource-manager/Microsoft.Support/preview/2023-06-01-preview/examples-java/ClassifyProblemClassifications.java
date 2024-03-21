
import com.azure.resourcemanager.support.models.ProblemClassificationsClassificationInput;

/**
 * Samples for ProblemClassificationsNoSubscription ClassifyProblems.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/support/resource-manager/Microsoft.Support/preview/2023-06-01-preview/examples/
     * ClassifyProblemClassifications.json
     */
    /**
     * Sample code: Classify list of problemClassifications for a specified Azure service.
     * 
     * @param manager Entry point to SupportManager.
     */
    public static void classifyListOfProblemClassificationsForASpecifiedAzureService(
        com.azure.resourcemanager.support.SupportManager manager) {
        manager.problemClassificationsNoSubscriptions().classifyProblemsWithResponse("serviceId1",
            new ProblemClassificationsClassificationInput().withIssueSummary("Can not connect to Windows VM"),
            com.azure.core.util.Context.NONE);
    }
}
