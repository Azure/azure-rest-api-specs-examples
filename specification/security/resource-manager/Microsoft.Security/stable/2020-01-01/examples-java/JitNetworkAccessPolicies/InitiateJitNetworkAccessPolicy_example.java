import com.azure.core.util.Context;
import com.azure.resourcemanager.security.models.JitNetworkAccessPolicyInitiatePort;
import com.azure.resourcemanager.security.models.JitNetworkAccessPolicyInitiateRequest;
import com.azure.resourcemanager.security.models.JitNetworkAccessPolicyInitiateVirtualMachine;
import java.util.Arrays;

/** Samples for JitNetworkAccessPolicies Initiate. */
public final class Main {
    /*
     * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2020-01-01/examples/JitNetworkAccessPolicies/InitiateJitNetworkAccessPolicy_example.json
     */
    /**
     * Sample code: Initiate an action on a JIT network access policy.
     *
     * @param manager Entry point to SecurityManager.
     */
    public static void initiateAnActionOnAJITNetworkAccessPolicy(
        com.azure.resourcemanager.security.SecurityManager manager) {
        manager
            .jitNetworkAccessPolicies()
            .initiateWithResponse(
                "myRg1",
                "westeurope",
                "default",
                new JitNetworkAccessPolicyInitiateRequest()
                    .withVirtualMachines(
                        Arrays
                            .asList(
                                new JitNetworkAccessPolicyInitiateVirtualMachine()
                                    .withId(
                                        "/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/myRg1/providers/Microsoft.Compute/virtualMachines/vm1")
                                    .withPorts(
                                        Arrays
                                            .asList(
                                                new JitNetworkAccessPolicyInitiatePort()
                                                    .withNumber(3389)
                                                    .withAllowedSourceAddressPrefix("192.127.0.2")))))
                    .withJustification("testing a new version of the product"),
                Context.NONE);
    }
}
