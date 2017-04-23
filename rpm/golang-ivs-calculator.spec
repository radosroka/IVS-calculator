%global provider        github
%global provider_tld    com
%global project         radosroka
%global repo            IVS-calculator
# https://github.com/radosroka/IVS-calculator
%global provider_prefix %{provider}.%{provider_tld}/%{project}/%{repo}
%global import_path     %{provider_prefix}
%global commit          5826ca2d1fc2318cb1ffe0c15b0fd20689f56b7b
%global shortcommit     %(c=%{commit}; echo ${c:0:7})

Name:           golang-ivs-calculator
Version:        1.1
Release:        1%{?dist}
Summary:        Simple calculator written in golang
License:        GPLv3
URL:            https://%{provider_prefix}
Source0:        IVS-calculator-%{version}.tar.gz

# e.g. el6 has ppc64 arch without gcc-go, so EA tag is required
ExclusiveArch:  %{?go_arches:%{go_arches}}%{!?go_arches:%{ix86} x86_64 aarch64 %{arm}}
# If go_compiler is not set to 1, there is no virtual provide. Use golang instead.
BuildRequires:  golang
BuildRequires:  golang-godoc
BuildRequires:  git
BuildRequires:  gtk2-devel
BuildRequires:  pango-devel

Requires: libcanberra-gtk2
Requires: PackageKit-gtk3-module

%description
%{summary}

Provides:      golang(%{import_path}) = %{version}-%{release}

%prep
%setup -q -n %{repo}-%{version}

%build
make build

%install
mkdir -p %{buildroot}/%{_bindir}
install -m 755 %{_builddir}/%{repo}-%{version}/bin/ivs-calc %{buildroot}/%{_bindir}

#%%check??
%files
#%license LICENSE
%doc doc/*.html
%{_bindir}/ivs-calc

%changelog
* Sun Apr 23 2017 Michal Cyprian <mcyprian@redhat.com> - 1.1-1
- Initial package
