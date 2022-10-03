import * as aws from "@pulumi/aws";

export interface OrganizationArgs {
  /**
   * The organization ID to import the Organization in the stack. If not set a new AWS Organization will be created. Defaults to undefined.
   */
  organizationId?: string;

  /**
   * The Organization policies to be applied.
   */
  policies?: OrganizationPoliciesArgs;

  /**
   * The list of AWS Account to be configured in the Organization.
   */
  accounts?: OrganizationAccountArgs[];
}

export interface OrganizationPoliciesArgs {
  /**
   * Deny IAM Account to leave the organization. Enabled by default.
   */
  denyLeaveOrganization?: OrganizationPolicyArgs;
}

export interface OrganizationPolicyArgs {
  /**
   * Enable the policy/
   */
  enabled?: boolean;

  /**
   * Import the policy with the given id
   */
  policyId?: string;
}

export interface OrganizationAccountArgs {
  /**
   * The name of the IAM Account.
   */
  name: string;

  /**
   * The email associated to the IAM Account.
   */
  email: string;

  /**
   * Admin role for the IAM Account.
   */
  adminRoleName?: string;

  /*
   *  The Organizational Unit to be used for the account.
   */
  ou?: string;

  /**
   * The AWS Account ID to be used to import the Account in the Organization. If not set, a new AWS Account will be created.
   */
  accountId?: string;

  /**
   * The parentId of the imported account.
   */
  parentId?: string;
}

export const defaultOrganizationAccount = {
  adminRoleName: "superadmin",
};

export interface PolicyData {
  name: string;
  description: string;
  policy: any;
}

export interface OrganizationPoliciesData {
  data: { [key: string]: PolicyData };
}

export const organizationPoliciesData: OrganizationPoliciesData = {
  data: {
    denyLeaveOrganization: {
      name: "Deny Leave Organization",
      description: "Prevent member accounts from leaving the organization",
      policy: {
        sid: "DenyLeaveOrganization",
        effect: "Deny",
        actions: ["organizations:LeaveOrganization"],
        resources: ["*"],
      },
    },
  },
};

export const defaultPolicies = {
  denyLeaveOrganization: {
    enabled: true,
  },
};

export const defaultOrganizationArgs = {
  allowedRegions: [],
  policies: defaultPolicies,
  accounts: [],
};

export interface OrganizationalUnitMapping {
  accountName: string;
  organizationalUnit: aws.organizations.OrganizationalUnit;
}
export interface AccountMappingArgs {
  accountName: string;
  account: aws.organizations.Account;
}