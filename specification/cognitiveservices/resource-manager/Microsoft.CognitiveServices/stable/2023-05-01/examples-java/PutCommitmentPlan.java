import com.azure.resourcemanager.cognitiveservices.fluent.models.CommitmentPlanInner;
import com.azure.resourcemanager.cognitiveservices.models.CommitmentPeriod;
import com.azure.resourcemanager.cognitiveservices.models.CommitmentPlanProperties;
import com.azure.resourcemanager.cognitiveservices.models.HostingModel;

/** Samples for CommitmentPlans CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2023-05-01/examples/PutCommitmentPlan.json
     */
    /**
     * Sample code: PutCommitmentPlan.
     *
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void putCommitmentPlan(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager
            .commitmentPlans()
            .createOrUpdateWithResponse(
                "resourceGroupName",
                "accountName",
                "commitmentPlanName",
                new CommitmentPlanInner()
                    .withProperties(
                        new CommitmentPlanProperties()
                            .withHostingModel(HostingModel.WEB)
                            .withPlanType("Speech2Text")
                            .withCurrent(new CommitmentPeriod().withTier("T1"))
                            .withAutoRenew(true)),
                com.azure.core.util.Context.NONE);
    }
}
