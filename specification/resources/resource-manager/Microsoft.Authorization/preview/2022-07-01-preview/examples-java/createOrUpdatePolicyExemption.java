
import com.azure.core.management.serializer.SerializerFactory;
import com.azure.core.util.serializer.SerializerEncoding;
import com.azure.resourcemanager.resources.fluent.models.PolicyExemptionInner;
import com.azure.resourcemanager.resources.models.ExemptionCategory;
import java.io.IOException;
import java.util.Arrays;

/**
 * Samples for PolicyExemptions CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/resources/resource-manager/Microsoft.Authorization/preview/2022-07-01-preview/examples/
     * createOrUpdatePolicyExemption.json
     */
    /**
     * Sample code: Create or update a policy exemption.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createOrUpdateAPolicyExemption(com.azure.resourcemanager.AzureResourceManager azure)
        throws IOException {
        azure.genericResources().manager().policyClient().getPolicyExemptions().createOrUpdateWithResponse(
            "subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/resourceGroups/demoCluster", "DemoExpensiveVM",
            new PolicyExemptionInner().withPolicyAssignmentId(
                "/subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/providers/Microsoft.Authorization/policyAssignments/CostManagement")
                .withPolicyDefinitionReferenceIds(Arrays.asList("Limit_Skus"))
                .withExemptionCategory(ExemptionCategory.WAIVER).withDisplayName("Exempt demo cluster")
                .withDescription("Exempt demo cluster from limit sku")
                .withMetadata(SerializerFactory.createDefaultManagementSerializerAdapter().deserialize(
                    "{\"reason\":\"Temporary exemption for a expensive VM demo\"}", Object.class,
                    SerializerEncoding.JSON)),
            com.azure.core.util.Context.NONE);
    }
}
