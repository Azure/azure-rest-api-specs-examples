
import com.azure.resourcemanager.devtestlabs.models.EvaluatePoliciesProperties;
import com.azure.resourcemanager.devtestlabs.models.EvaluatePoliciesRequest;
import java.util.Arrays;

/**
 * Samples for PolicySets EvaluatePolicies.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/
     * PolicySets_EvaluatePolicies.json
     */
    /**
     * Sample code: PolicySets_EvaluatePolicies.
     * 
     * @param manager Entry point to DevTestLabsManager.
     */
    public static void policySetsEvaluatePolicies(com.azure.resourcemanager.devtestlabs.DevTestLabsManager manager) {
        manager.policySets().evaluatePoliciesWithResponse("resourceGroupName", "{labName}", "{policySetName}",
            new EvaluatePoliciesRequest().withPolicies(
                Arrays.asList(new EvaluatePoliciesProperties().withFactName("LabVmCount").withValueOffset("1"))),
            com.azure.core.util.Context.NONE);
    }
}
