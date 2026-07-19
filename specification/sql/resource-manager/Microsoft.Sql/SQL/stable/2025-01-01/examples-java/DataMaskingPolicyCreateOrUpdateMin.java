
import com.azure.resourcemanager.sql.fluent.models.DataMaskingPolicyInner;
import com.azure.resourcemanager.sql.models.DataMaskingPolicyName;
import com.azure.resourcemanager.sql.models.DataMaskingState;

/**
 * Samples for DataMaskingPolicies CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/DataMaskingPolicyCreateOrUpdateMin.json
     */
    /**
     * Sample code: Create or update data masking policy min.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void createOrUpdateDataMaskingPolicyMin(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getDataMaskingPolicies().createOrUpdateWithResponse("sqlcrudtest-6852",
            "sqlcrudtest-2080", "sqlcrudtest-331", DataMaskingPolicyName.DEFAULT,
            new DataMaskingPolicyInner().withDataMaskingState(DataMaskingState.ENABLED),
            com.azure.core.util.Context.NONE);
    }
}
