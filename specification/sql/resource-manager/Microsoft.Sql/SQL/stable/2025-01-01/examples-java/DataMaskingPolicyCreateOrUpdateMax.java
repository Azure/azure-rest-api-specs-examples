
import com.azure.resourcemanager.sql.fluent.models.DataMaskingPolicyInner;
import com.azure.resourcemanager.sql.models.DataMaskingPolicyName;
import com.azure.resourcemanager.sql.models.DataMaskingState;

/**
 * Samples for DataMaskingPolicies CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/DataMaskingPolicyCreateOrUpdateMax.json
     */
    /**
     * Sample code: Create or update data masking policy max.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void createOrUpdateDataMaskingPolicyMax(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getDataMaskingPolicies().createOrUpdateWithResponse("sqlcrudtest-6852",
            "sqlcrudtest-2080", "sqlcrudtest-331", DataMaskingPolicyName.DEFAULT, new DataMaskingPolicyInner()
                .withDataMaskingState(DataMaskingState.ENABLED).withExemptPrincipals("testuser;"),
            com.azure.core.util.Context.NONE);
    }
}
