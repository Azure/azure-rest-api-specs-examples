import com.azure.core.util.Context;
import com.azure.resourcemanager.policyinsights.fluent.models.AttestationInner;
import com.azure.resourcemanager.policyinsights.models.AttestationEvidence;
import com.azure.resourcemanager.policyinsights.models.ComplianceState;
import java.time.OffsetDateTime;
import java.util.Arrays;

/** Samples for Attestations CreateOrUpdateAtResource. */
public final class Main {
    /*
     * x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2021-01-01/examples/Attestations_CreateResourceScope.json
     */
    /**
     * Sample code: Create attestation at individual resource scope.
     *
     * @param manager Entry point to PolicyInsightsManager.
     */
    public static void createAttestationAtIndividualResourceScope(
        com.azure.resourcemanager.policyinsights.PolicyInsightsManager manager) {
        manager
            .attestations()
            .createOrUpdateAtResource(
                "subscriptions/35ee058e-5fa0-414c-8145-3ebb8d09b6e2/resourcegroups/myrg/providers/microsoft.compute/virtualMachines/devVM",
                "790996e6-9871-4b1f-9cd9-ec42cd6ced1e",
                new AttestationInner()
                    .withPolicyAssignmentId(
                        "/subscriptions/35ee058e-5fa0-414c-8145-3ebb8d09b6e2/providers/microsoft.authorization/policyassignments/b101830944f246d8a14088c5")
                    .withPolicyDefinitionReferenceId("0b158b46-ff42-4799-8e39-08a5c23b4551")
                    .withComplianceState(ComplianceState.COMPLIANT)
                    .withExpiresOn(OffsetDateTime.parse("2021-06-15T00:00:00Z"))
                    .withOwner("55a32e28-3aa5-4eea-9b5a-4cd85153b966")
                    .withComments("This subscription has passed a security audit.")
                    .withEvidence(
                        Arrays
                            .asList(
                                new AttestationEvidence()
                                    .withDescription("The results of the security audit.")
                                    .withSourceUri(
                                        "https://gist.github.com/contoso/9573e238762c60166c090ae16b814011"))),
                Context.NONE);
    }
}
